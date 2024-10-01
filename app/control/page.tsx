"use client";
import QuestionControl from "@/components/question-control";
import RemovePlayer from "@/components/remove-player";
import UpdateScore from "@/components/update-score";
import { TextField } from "@mui/material";
import { useRef, useState, useEffect } from "react";
import { Player } from "@/types";
import Leaderboard from "@/components/leaderboard";
import BuzzedIn from "@/components/buzzed-in";


export default function ControlPage() {
  const [password, setPassword] = useState("")
  const [errorMessage, setErrorMessage] = useState("")
  const [players, setPlayers] = useState<Player[]>([])
  const wsRef = useRef<WebSocket | null>(null);
  const [waitingToReconnect, setWaitingToReconnect] = useState(false);

  const resetError = () => {
    setErrorMessage("")
  }

  const fetchAuthEndpoint = (endpoint: string, method?: string) => {
    if (password === "") {
      setErrorMessage("password is required")
      return
    }
    const pw = (endpoint.includes("?") ? "&" : "?") + `password=${password}`
    fetch(`/api/auth/${endpoint}${pw}`, {
      method: method ?? "GET"
    }).then(res => res.text()).then(msg => (
      msg != "success" ? setErrorMessage(msg) : resetError()
    ))
  }

  useEffect(() => {

    if (waitingToReconnect) {
      return;
    }

    // Only set up the websocket once
    if (!wsRef.current) {
      const ws = new WebSocket("/api/players");
      wsRef.current = ws;

      ws.onerror = (e) => console.error(e);

      ws.onopen = () => {
        console.log('ws opened');
        // client.send('ping');
      };

      ws.onclose = () => {

        if (wsRef.current) {
          console.log('ws closed by server');
        } else {
          console.log('ws closed by app component unmount');
          return;
        }

        if (waitingToReconnect) {
          return;
        };

        console.log('ws closed');

        setWaitingToReconnect(true);

        setTimeout(() => setWaitingToReconnect(false), 1000);
      };

      ws.onmessage = (e) => {
        console.log("received", e);
        const d = JSON.parse(e.data) ?? []
        setPlayers(d);
        console.log(d)
      };


      return () => {
        console.log('Cleanup');
        ws.close();
        wsRef.current = null;
      }
    }

  }, [waitingToReconnect]);


  return (
    <main className="flex flex-col items-center p-2">
      <h1 className="text-4xl p-2 font-semibold m-4">Control Panel</h1>
      <div>
        <TextField variant="outlined" label="Password" type="password" value={password} required 
          onChange={(e)=>{
          resetError();
          setPassword(e.target.value);
          }} 
        />
        <p
          className="text-md text-center p-1"
          style={{
            visibility: errorMessage.length > 0 ? "visible" : "hidden",
            color: "red"
          }}
        >{errorMessage}</p>
      </div>
      <div className="flex flex-col sm:flex-row">
        <UpdateScore submit={(name, delta) => (
          fetchAuthEndpoint(`score?name=${name}&amount=${delta}`, "PUT")
        )} 
          players={players}
        />
        <QuestionControl reset={() => {
          fetchAuthEndpoint(`reset`, "POST")
        }} />
        <RemovePlayer remove={(name) => (
          fetchAuthEndpoint(`player?name=${name}`, "DELETE")
        )}
          players={players}
        />
      </div>
      <div className="flex">
        <Leaderboard />
        <BuzzedIn />
      </div>
    </main>
  )
}
