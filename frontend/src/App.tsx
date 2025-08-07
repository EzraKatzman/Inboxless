import Navbar from "./components/Layout/Navbar"
import CreateInbox from "./components/Buttons/CreateInbox"
import Hero from "./components/Layout/Hero"
import Faq from "./components/Layout/Faq"

export default function App() {

  return (
    <div className="bg-neutral-100 dark:bg-neutral-900">
      <main className="overflow-x-auto">
        <header className="sticky top-0 z-40 flex-none w-full bg-neutral-100 dark:bg-neutral-900">
          <Navbar />
        </header>
        <div className="overflow-x-auto container mx-auto max-w-6xl pb-8 text-center bg-neutral-100 dark:bg-neutral-900 text-neutral-900 dark:text-neutral-100">
          <h1 className="text-4xl lg:text-6xl font-bold my-8">Let's make you an email like </h1>
          <Hero />
          <CreateInbox />
          <Faq />
        </div>
      </main>
    </div>
  )
}
