import { useEffect, useState } from "react";
import {
    LightIcon,
    DarkIcon,
} from '../icons/themeIcons';

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
            className="p-2 rounded-full bg-neutral-200 hover:bg-neutral-300 text-neutral-900 hover:text-violet-700 cursor-pointer"
            aria-label="Toggle theme"
        >
            {theme === "dark" ? <LightIcon className="w-5 h-5"/> : <DarkIcon className="w-5 h-5"/>}
        </button>
    )
}