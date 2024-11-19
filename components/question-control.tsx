"use client";
import { Button } from "@mui/material";

const QuestionControl = ({ reset }: { reset: () => void }) => {
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
                sx={{ margin: "0.25rem" }}
                onClick={() => reset()}
            >
                Reset Buzzers
            </Button>
        </div>
    );
};

export default QuestionControl;
