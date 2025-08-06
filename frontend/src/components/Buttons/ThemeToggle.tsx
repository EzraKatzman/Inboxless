import { useEffect, useState } from "react";
import {
    LightIcon,
    DarkIcon,
} from '../../../public/icons/themeIcons';

export default function ThemeToggle() {
    const getSystemTheme = () => {
        return window.matchMedia("(prefers-color-scheme: dark)").matches ? "dark" : "light";
    }

    const [theme, setTheme] = useState<"light" | "dark">(getSystemTheme());

    useEffect(() => {
        const root = document.body;
        root.classList.remove(theme === "light" ? "dark" : "light");
        root.classList.add(theme);  

        localStorage.setItem("theme", theme);
    }, [theme]);

    useEffect(() => {
        const savedTheme = localStorage.getItem("theme");
        if (savedTheme === "light" || savedTheme === "dark") {
            setTheme(savedTheme);
        } else {
            setTheme(getSystemTheme());
        }
    }, []);

    const toggleTheme = () => setTheme(theme === "dark" ? "light" : "dark");

    return (
        <button
            onClick={toggleTheme}
            className="p-2 rounded-full bg-gray-200 hover:bg-gray-300 text-gray-900"
            aria-label="Toggle theme"
        >
            {theme === "dark" ? <LightIcon className="w-6 h-6"/> : <DarkIcon className="w-6 h-6"/>}
        </button>
    )
}