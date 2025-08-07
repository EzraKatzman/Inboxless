import { Typewriter } from "react-simple-typewriter"

export default function Hero() {
    return (
        <>
        <span className="text-4xl lg:text-6xl font-bold my-8 text-violet-700 dark:text-violet-600">
            <Typewriter
                words={["inboxname 1", "inboxname 2", "inboxname 3"]}
                loop={true}
                cursor={false}
                cursorStyle="|"
                typeSpeed={70}
                deleteSpeed={60}  
                delaySpeed={1000}
            />
            @inboxless.io
        </span>
        </>
    );
}