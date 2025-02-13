import { SignInButton } from "@clerk/clerk-react";


function Landing(){
    return (
        <div>
            <h1>Welcome to the Aux Sesh</h1>
            <p>Sign in to access the app</p>
            <SignInButton />
        </div>
    )
}

export default Landing;