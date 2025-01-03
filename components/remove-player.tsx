"use client";
import { useState } from "react";
import { Button, MenuItem, Modal, Select } from "@mui/material";
import { Player } from "@/types";

const RemovePlayer = ({
    players,
    remove,
}: {
    players: Player[];
    remove: (token: string) => void;
}) => {
    const [name, setName] = useState("");
    const [modalOpen, setModalOpen] = useState(false);
    const closeModal = () => setModalOpen(false);

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
                    onClick={() => setModalOpen(true)}
                >
                    Remove
                </Button>
            </div>
            <Modal open={modalOpen} onClose={closeModal}>
                <div
                    className="absolute top-1/2 left-1/2 transform -translate-x-1/2 -translate-y-1/2 
		    w-60 aspect-square bg-white rounded-3xl flex flex-col items-center justify-center"
                >
                    <h5 className="text-xl p-1 m-2 text-center">
                        Delete player:{" "}
                        <span className="text-[#F00] font-semibold">
                            {name}
                        </span>
                        ? This action is permanent.
                    </h5>
                    <div className="w-[60%] flex justify-between">
                        <Button variant="outlined" onClick={closeModal}>
                            No
                        </Button>
                        <Button
                            variant="contained"
                            onClick={() => {
                                remove(name);
                                setName("");
                                setModalOpen(false);
                            }}
                        >
                            Yes
                        </Button>
                    </div>
                </div>
            </Modal>
        </div>
    );
};

export default RemovePlayer;
