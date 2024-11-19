"use client";
import { useState } from "react";
import { Button, Modal } from "@mui/material";

const ResetGame = ({ submit }: { submit: () => void }) => {
    const [modalOpen, setModalOpen] = useState(false);
    const closeModal = () => setModalOpen(false);

    return (
        <div>
            <Button
                variant="contained"
                sx={{ margin: "0.25rem" }}
                onClick={() => setModalOpen(true)}
            >
                Reset Game
            </Button>
            <Modal open={modalOpen} onClose={closeModal}>
                <div
                    className="absolute top-1/2 left-1/2 transform -translate-x-1/2 -translate-y-1/2 
		    w-72 p-4 aspect-square bg-white rounded-3xl flex flex-col items-center justify-center"
                >
                    <h5 className="text-xl text-center p-1 m-2">
                        Are you sure you want to reset the game? This action is
                        permanent.
                    </h5>
                    <div className="w-[60%] flex justify-between">
                        <Button variant="outlined" onClick={closeModal}>
                            No
                        </Button>
                        <Button
                            variant="contained"
                            onClick={() => {
                                submit();
                                closeModal();
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

export default ResetGame;
