import { useState } from "react";
import {RightArrowIcon} from "../../../public/icons/rightArrow";
import { createInbox } from "../../api/index";

export default function CreateInbox() {
    const [loading, setLoading] = useState(false);
    const [result, setResult] = useState<string | null>(null);

    const handleCreateInbox = async () => {
        setLoading(true);
        try {
            const response = await createInbox();
            setResult(response.id); // Assuming the response contains an 'id' field
        } catch (err: any) {
            setResult(err.message || "Error creating inbox");
        } finally {
            setLoading(false);
        }
    };

    return (
        <>
            <button 
                className="inline-flex items-center justify-center text-center font-medium text-white bg-blue-700 hover:bg-blue-800 dark:bg-blue-600 dark:hover:bg-blue-700 px-6 py-3.5 rounded-lg"
                onClick={handleCreateInbox}
                disabled={loading}
                aria-label="Create Inbox"
                >
                {loading ? "Creating Inbox..." : "Create Inbox"}
                <RightArrowIcon className="ml-2 -mr-1"/>
            </button>
            {result && (
                <div className="mt-4 text-sm text-gray-700 dark:text-gray-300">
                    {result}
                </div>
            )}
        </>
    );
}