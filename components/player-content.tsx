"use client";
import { useEffect, useRef, useState } from "react";
import { useRouter, useParams } from "next/navigation";

const PlayerContent = () => {
    const [buttonReady, setButtonReady] = useState(false);

    const router = useRouter();
    const name = decodeURI(useParams()["player-name"] as string);
    const wsRef = useRef<WebSocket | null>(null);

    //    useEffect(() => {
    // if (ws leaderboard === null) {
    //     setWs(new WebSocket("/api/buzz"))
    //     return
    // }
    // ws.onopen = () => {
    //     console.log("connected to ws")
    //     ws.send(name);
    // }
    // ws.onmessage = (e) => {
    //     console.log(e.data)
    //     if (e.data == "ready") {
    // 	setButtonReady(true)
    //     } else {
    // 	setButtonReady(false)
    //     }
    // }
    // ws.onclose = () => {
    //     console.log("disconnected from server");
    //     router.push("/");
    // }
    //    }, [ws])

    useEffect(() => {
        if (!wsRef.current) {
            console.log("connecting to websocke");
            const ws = new WebSocket("/api/buzz");
            wsRef.current = ws;

            ws.onerror = (e) => console.error(e);

            ws.onopen = () => {
                console.log("ws buzz opened, writing");
                ws.send(name);
            };

            ws.onclose = () => {
                console.log("disconnected from server");
                router.push("/");
            };

            ws.onmessage = (e) => {
                console.log(e.data);
                setButtonReady(e.data == "ready");
            };

            // return () => {
            //   console.log("cleanup ws");
            //   ws.close();
            //   wsRef.current = null;
            // };
        }
    }, []);

    return (
        <main className="flex flex-col items-center justify-center h-screen bg-blue-200">
            <h1 className="p-2 m-4 text-4xl">
                Welcome{" "}
                <span className="underline" id="name">
                    {name}
                </span>
            </h1>
            <button
                onClick={() => {
                    try {
                        wsRef.current?.send(name);
                        setButtonReady(false);
                    } catch (e) {
                        console.error(e);
                    }
                }}
                disabled={!buttonReady}
                style={{
                    backgroundColor: buttonReady ? "lightgreen" : "gray",
                    color: buttonReady ? "black" : "darkgray",
                }}
                className="p-6 w-56 aspect-square select-none cursor-pointer flex flex-col justify-center items-center text-center text-6xl rounded-full transition-all duration-75"
            >
                Buzz
            </button>
        </main>
    );
};

export default PlayerContent;
