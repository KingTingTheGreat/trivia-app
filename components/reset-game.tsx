import { Button } from "@mui/material";

const ResetGame = ({ submit }: { submit: () => void }) => {
    return (
        <div>
            <Button variant="contained" onClick={submit}>
                Reset Game
            </Button>
        </div>
    );
};

export default ResetGame;
