import Navbar from "./components/Miscellaneous/Navbar"
import CreateInbox from "./components/Buttons/CreateInbox"

export default function App() {

  return (
    <div className="bg-white dark:bg-gray-900">
      <main className="overflow-x-auto">
        <header className="sticky top-0 z-40 flex-none w-full bg-white dark:bg-gray-900">
          <Navbar />
        </header>
        <div className="overflow-x-auto container mx-auto max-w-6xl pb-8 text-center bg-white dark:bg-gray-900 text-black dark:text-white">
          <h1 className="text-4xl lg:text-6xl font-bold my-8">Welcome to Inboxless</h1>
          <CreateInbox />
        </div>
      </main>
    </div>
  )
}
