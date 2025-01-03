"use client";
import { Player } from "@/types";
import { Button, MenuItem, Select, TextField } from "@mui/material";
import { useState } from "react";

const UpdateScore = ({
    players,
    update,
    clear,
}: {
    players: Player[];
    update: (name: string, delta: string) => void;
    clear: (name: string) => void;
}) => {
    const [name, setName] = useState("");
    const [delta, setDelta] = useState("0");

    return (
        <div className="m-1">
            <h4 className="p-1 text-xl text-center">Update Score</h4>
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
                <TextField
                    variant="outlined"
                    label="Δ"
                    type="number"
                    sx={{ width: "5rem" }}
                    value={delta}
                    onChange={(e) => {
                        setDelta(e.target.value);
                    }}
                />
                <Button
                    variant="contained"
                    sx={{ padding: "0.75rem", margin: "0.5rem" }}
                    onClick={() => {
                        update(name, delta);
                        setName("");
                        setDelta("0");
                    }}
                >
                    Add to score
                </Button>
                <Button
                    variant="contained"
                    sx={{ padding: "0.75rem", margin: "0.5rem" }}
                    onClick={() => {
                        clear(name);
                        setName("");
                    }}
                >
                    Clear score
                </Button>
            </div>
        </div>
    );
};

export default UpdateScore;
