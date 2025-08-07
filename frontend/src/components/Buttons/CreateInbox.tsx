import { useState } from "react";
import {RightArrowIcon} from "../icons/rightArrow";
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
        <div className="mt-8">
            <button 
                className="inline-flex items-center justify-center text-center font-semibold text-neutral-100 bg-violet-700 hover:bg-violet-800 dark:bg-violet-600 dark:hover:bg-violet-700 px-6 py-3.5 rounded-lg"
                onClick={handleCreateInbox}
                disabled={loading}
                aria-label="Create Inbox"
                >
                {loading ? "Creating Inbox..." : "Create Inbox"}
                <RightArrowIcon className="ml-2 -mr-1"/>
            </button>
            {result && (
                <div className="mt-4 text-sm text-neutral-700 dark:text-neutral-300">
                    {result}
                </div>
            )}
        </div>
    );
}