
import { SignedIn, SignedOut, SignInButton } from "@clerk/clerk-react";
import Home from "./pages/Home";

export default function App() {
  return (
    <header>
      <SignedOut>
        <SignInButton />
      </SignedOut>
      <SignedIn>
       <Home/>
      </SignedIn>
    </header>
  );
}