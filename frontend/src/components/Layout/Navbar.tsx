import ThemeToggle from "../Buttons/ThemeToggle";

export default function Navbar() {
  return (
    <nav className="max-w-8xl m-auto px-4 py-2.5 text-center bg-neutral-100 dark:bg-neutral-900 text-neutral-900 dark:text-neutral-100">
      <div className="mx-auto flex flex-wrap justify-between items-center container">
        <div className="w-48 flex-shrink-0 text-4xl font-bold text-neutral-800 dark:text-neutral-100">
          INBOXLESS
        </div>
        <div className="flex-1 flex justify-center">
          <div className="space-x-4">
            <a
              href="/"
              className="text-neutral-700 hover:text-violet-700 dark:text-neutral-400 dark:hover:text-neutral-100"
            >
              Home
            </a>
            <a
              href="#faq"
              className="text-neutral-700 hover:text-violet-700 dark:text-neutral-400 dark:hover:text-neutral-100"
            >
              FAQ
            </a>
            <a
              href="/about"
              className="text-neutral-700 hover:text-violet-700 dark:text-neutral-400 dark:hover:text-neutral-100" 
            >
              About Us
            </a>
          </div>
        </div>
        <div className="w-48 flex-shrink-0">
          <ThemeToggle />
        </div>
      </div>
    </nav>
  );
}