"use client";
import { useRef, useState, useEffect } from "react";
import GameContent, { TableData, TableRow } from "./game-content";
import { Player } from "@/types";

const BuzzedIn = () => {
    const [players, setPlayers] = useState<Player[]>([]);
    const wsRef = useRef<WebSocket | null>(null);
    const [waitingToReconnect, setWaitingToReconnect] = useState(false);
    const buzzer = useRef<HTMLAudioElement | null>(null);

    useEffect(() => {
        buzzer.current = new Audio("/buzzer.mp3");
    }, []);

    useEffect(() => {
        if (waitingToReconnect) {
            return;
        }

        // Only set up the websocket once
        if (!wsRef.current) {
            const ws = new WebSocket("/api/buzzed-in");
            wsRef.current = ws;

            ws.onerror = (e) => console.error(e);

            ws.onopen = () => {
                console.log("ws buzzed in opened");
                // client.send('ping');
            };

            ws.onclose = () => {
                if (wsRef.current) {
                    console.log("ws buzzed in closed by server");
                } else {
                    console.log("ws buzzed in closed by app component unmount");
                    return;
                }

                if (waitingToReconnect) {
                    return;
                }

                console.log("ws buzzed in closed");

                setWaitingToReconnect(true);

                setTimeout(() => setWaitingToReconnect(false), 100);
            };

            ws.onmessage = (e) => {
                console.log("ws buzzed in received", e);
                const d = JSON.parse(e.data) ?? [];
                setPlayers(d);
                console.log(d);
            };

            return () => {
                console.log("Cleanup buzzed in ws");
                ws.close();
                wsRef.current = null;
            };
        }
    }, [waitingToReconnect]);

    useEffect(() => {
        if (players.length > 0) {
            buzzer.current?.play();
        }
    }, [players]);

    const mapFunc = (player: Player, index: number): React.ReactNode => (
        <TableRow index={index} key={player.Name + player.Time}>
            <TableData>{player.Name}</TableData>
            <TableData>{player.Time}</TableData>
        </TableRow>
    );

    return wsRef.current ? (
        <GameContent
            title="Buzzed In"
            headers={["Name", "Time"]}
            content={players}
            mapFunc={mapFunc}
        />
    ) : (
        <p>Not connected to server</p>
    );
};

export default BuzzedIn;
