import dynamic from "next/dynamic";

export default function PlayerPage() {
    const NoSsrPlayerContent = dynamic(() => import("@/components/player-content"), {
	ssr: false,
    })
    return <NoSsrPlayerContent />
}
