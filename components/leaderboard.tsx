"use client";
import { useState, useEffect, useRef } from "react";
import GameContent, { TableData, TableRow } from "./game-content";
import { Player } from "@/types";

const Leaderboard = () => {
    const [players, setPlayers] = useState<Player[]>([]);
    const wsRef = useRef<WebSocket | null>(null);
    const [waitingToReconnect, setWaitingToReconnect] = useState(false);

    //    useEffect(() => {
    // if (ws === null) {
    //     setWs(new WebSocket("/api/leaderboard"));
    //     return;
    // }
    // ws.onopen = () => {
    //     console.log("connected to leaderboard");
    // }
    // ws.onmessage = (e) => {
    //     console.log("leaderboard", e.data);
    //     setPlayers(JSON.parse(e.data) ?? []);
    // }
    // ws.onclose = () => {
    //     console.log("disconnected from leaderboard");
    //     setPlayers([]);
    //     setTimeout(() => {
    // 	setWs(new WebSocket("/api/leaderboard"));
    //     }, 100)
    // }
    //    }, [ws])

    useEffect(() => {
        if (waitingToReconnect) {
            return;
        }

        // Only set up the websocket once
        if (!wsRef.current) {
            const ws = new WebSocket("/api/leaderboard");
            wsRef.current = ws;

            ws.onerror = (e) => console.error(e);

            ws.onopen = () => {
                console.log("ws leaderboard opened");
                // client.send('ping');
            };

            ws.onclose = () => {
                if (wsRef.current) {
                    console.log("ws leaderboard closed by server");
                } else {
                    console.log(
                        "ws leaderboard closed by app component unmount"
                    );
                    return;
                }

                if (waitingToReconnect) {
                    return;
                }

                console.log("ws leaderboard closed");

                setWaitingToReconnect(true);

                setTimeout(() => setWaitingToReconnect(false), 500);
            };

            ws.onmessage = (e) => {
                console.log("ws leaderboard received", e);
                const d = JSON.parse(e.data) ?? [];
                setPlayers(d);
                console.log(d);
            };

            return () => {
                console.log("Cleanup ws leaderboard");
                ws.close();
                wsRef.current = null;
            };
        }
    }, [waitingToReconnect]);

    const mapFunc = (player: Player, index: number): React.ReactNode => (
        <TableRow index={index} key={player.Name + player.Score}>
            <TableData>{player.Name}</TableData>
            <TableData>{player.Score}</TableData>
        </TableRow>
    );

    return wsRef.current ? (
        <GameContent
            title="Leaderboard"
            headers={["Name", "Score"]}
            content={players}
            mapFunc={mapFunc}
        />
    ) : (
        <p>Not connected to server</p>
    );
};

export default Leaderboard;
