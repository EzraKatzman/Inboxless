import ThemeToggle from "../Buttons/ThemeToggle";

export default function Navbar() {
  return (
    <nav className="max-w-8xl m-auto px-4 py-2.5 text-center bg-white dark:bg-gray-900 text-black dark:text-white">
      <div className="mx-auto flex flex-wrap justify-between items-center container">
        <div className="w-48 flex-shrink-0 text-4xl font-bold text-gray-800 dark:text-white">
          INBOXLESS
        </div>
        <div className="flex-1 flex justify-center">
          <div className="space-x-4">
            <a
              href="/"
              className="text-gray-700 hover:text-black dark:text-gray-300 dark:hover:text-white"
            >
              Home
            </a>
            <a
              href="#faq"
              className="text-gray-700 hover:text-black dark:text-gray-300 dark:hover:text-white"
            >
              FAQ
            </a>
            <a
              href="/about"
              className="text-gray-700 hover:text-black dark:text-gray-300 dark:hover:text-white" 
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