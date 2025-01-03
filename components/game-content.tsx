import { CSSProperties } from "react";
import { Player } from "@/types";
import { CircularProgress } from "@mui/material";

export const TableRow = ({
    children,
    index,
}: {
    children: React.ReactNode;
    index: number;
}) => (
    <tr className="border-2" key={"row: " + index}>
        {children}
    </tr>
);

export const TableData = ({
    children,
    style,
}: {
    children: React.ReactNode;
    style?: CSSProperties;
}) => (
    <td
        className="text-center py-2 px-4 border-2 text-lg font-medium"
        style={style}
    >
        {children}
    </td>
);

const GameContent = ({
    title,
    headers,
    content,
    mapFunc,
    loading,
}: {
    title: string;
    headers: string[];
    content: Player[];
    mapFunc: (row: Player, index: number) => React.ReactNode;
    loading: boolean;
}) => (
    <div className="flex flex-col items-center">
        <h2 className="text-5xl max-sm:text-3xl font-semibold p-2 m-1">
            {title}
        </h2>
        {!loading ? (
            <table className="border-collapse">
                <thead>
                    <tr>
                        {headers.map((header, index) => (
                            <th
                                className="px-6 py-2 border-solid border-2 text-xl"
                                key={"header: " + index}
                            >
                                {header}
                            </th>
                        ))}
                    </tr>
                </thead>
                <tbody>
                    {content.map((row, index) => mapFunc(row, index))}
                </tbody>
            </table>
        ) : (
            <p>
                Connecting to server... <CircularProgress />
            </p>
        )}
    </div>
);

export default GameContent;
