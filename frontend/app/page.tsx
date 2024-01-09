import { Suspense } from "react"
import { Spinner } from "@chakra-ui/react"
import QuizLauncher from "./components/QuizLauncher"

export default function Home() {
  return (
    <main>
      
      <div className="m-10 text-center">
      <Suspense>
        <QuizLauncher/>
      </Suspense>
      </div>
    </main>
  )
}