"use client";
import { useState } from "react";
import { Button, Modal } from "@mui/material";

const QuestionControl = ({ reset }: { reset: () => void }) => {
    const [modalOpen, setModalOpen] = useState(false);
    const closeModal = () => setModalOpen(false);

    return (
        <div className="flex flex-col items-center m-2 p-3">
            <h4 className="p-1 text-xl text-center">Question Control</h4>
            {/*<div className="flex w-40 justify-around p-1 m-1">
                <Button
                    variant="outlined"
                    onClick={() => {
                        console.log("prev");
                    }}
                >
                    Prev
                </Button>
                <Button
                    variant="outlined"
                    onClick={() => {
                        console.log("next");
                    }}
                >
                    Next
                </Button>
            </div>*/}
            <Button
                variant="contained"
                className="m-1"
                onClick={() => setModalOpen(true)}
            >
                Reset Buzzers
            </Button>
            <Modal open={modalOpen} onClose={closeModal}>
                <div
                    className="absolute top-1/2 left-1/2 transform -translate-x-1/2 -translate-y-1/2 
		    w-60 aspect-square bg-white rounded-3xl flex flex-col items-center justify-center"
                >
                    <h5 className="text-xl text-center p-1 m-2">
                        Reset buzzers?
                    </h5>
                    <div className="w-[60%] flex justify-between">
                        <Button variant="outlined" onClick={closeModal}>
                            No
                        </Button>
                        <Button
                            variant="contained"
                            onClick={() => {
                                reset();
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

export default QuestionControl;
