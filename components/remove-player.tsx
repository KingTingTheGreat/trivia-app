"use client";
import { useState } from "react";
import { Button, MenuItem, Select } from "@mui/material";
import { Player } from "@/types";

const RemovePlayer = ({
    players,
    remove,
}: {
    players: Player[];
    remove: (token: string) => void;
}) => {
    const [name, setName] = useState("");

    return (
        <div className="m-1">
            <h4 className="p-1 text-xl text-center">Remove Player</h4>
            <div className="flex items-center">
                <Select
                    value={name}
                    label="Player"
                    sx={{ minWidth: "8rem" }}
                    onChange={(e) => setName(e.target.value)}
                >
                    <MenuItem value=""></MenuItem>
                    {players.map((p) => (
                        <MenuItem value={p.Name} key={p.Token}>
                            {p.Name}
                        </MenuItem>
                    ))}
                </Select>
                <Button
                    variant="contained"
                    sx={{ padding: "0.75rem", margin: "0.5rem" }}
                    onClick={() => remove(name)}
                >
                    Remove
                </Button>
            </div>
        </div>
    );
};

export default RemovePlayer;
