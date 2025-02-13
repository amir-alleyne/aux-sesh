
import { SignInButton, SignedIn, SignedOut } from "@clerk/clerk-react";
import Home from "./pages/Home";
import Landing from "./pages/Landing";

export default function App() {
  return (
    <header>
      <SignedOut>
        <Landing/>
      </SignedOut>
      <SignedIn>
       <Home/>
      </SignedIn>
    </header>
  );
}