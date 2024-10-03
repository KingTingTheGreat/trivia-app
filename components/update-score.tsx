"use client";
import { Player } from "@/types";
import { Button, MenuItem, Select, TextField } from "@mui/material";
import { useState } from "react";

const UpdateScore = ({
    players,
    submit,
}: {
    players: Player[];
    submit: (name: string, delta: string) => void;
}) => {
    const [name, setName] = useState("");
    const [delta, setDelta] = useState("0");

    return (
        <div className="p-2 m-4">
            <h4 className="p-1 text-xl text-center">Update Player Score</h4>
            <div className="flex items-center">
                <Select
                    value={name}
                    label="Player"
                    className="min-w-32"
                    onChange={(e) => setName(e.target.value)}
                >
                    <MenuItem value=""></MenuItem>
                    {players.map((p) => (
                        <MenuItem value={p.Name} key={p.Token}>
                            {p.Name}
                        </MenuItem>
                    ))}
                </Select>
                <TextField
                    variant="outlined"
                    label="Δ"
                    type="number"
                    className="w-20"
                    value={delta}
                    onChange={(e) => {
                        setDelta(e.target.value);
                    }}
                />
                <Button
                    variant="contained"
                    className="p-3 m-2"
                    onClick={() => {
                        submit(name, delta);
                        setName("");
                        setDelta("0");
                    }}
                >
                    Update
                </Button>
            </div>
        </div>
    );
};

export default UpdateScore;
