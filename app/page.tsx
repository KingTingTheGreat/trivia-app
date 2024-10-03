import dynamic from "next/dynamic";

export default function HomePage() {
    const NoSsrHomeForm = dynamic(() => import("@/components/home-form"), {
        ssr: false,
    });
    return <NoSsrHomeForm />;
}
