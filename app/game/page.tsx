import Leaderboard from "@/components/leaderboard";
import BuzzedIn from "@/components/buzzed-in";

export default function GamePage() {
    return (
        <div className="flex justify-around">
            <Leaderboard />
            <BuzzedIn />
        </div>
    );
}
