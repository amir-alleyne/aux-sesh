import { UserButton } from "@clerk/clerk-react";
import Button from '@mui/material/Button';
import { useSpotifySignIn } from "../hooks/useSpotifySignIn";
import { useCreateSession } from "../hooks/useCreateSession";


function Home() {

  const { response, loading: signInLoading, error:signInError, signInWithSpotify } = useSpotifySignIn();
  const { session, loading: sessionLoading, error: sessionError, createAuxSession } = useCreateSession();
const handleSpotifySignIn = async () => {
    console.log('Sign In with Spotify');
    try {
        await signInWithSpotify();
    } catch (err) {
        console.error('Sign In failed:', err);
    }
    };

const handleCreateSession = async () => {
    try {
        await createAuxSession();
    } catch (err) {
        console.error('Create Session failed:', err);
    }
    }
  return (
    <>
      <div>
      <Button variant="outlined" onClick={handleSpotifySignIn}>Sign In With Spotify</Button>
    </div>
    <div>
      {signInLoading && <p>Loading...</p>}
      {signInError && <p>{signInError.message}</p>}
      {response && <p>{response.isSignedIn}</p>}
    </div>
    <div>
      <Button variant="outlined" onClick={handleCreateSession}>Create Session</Button>
    </div>
    <div>
      {sessionLoading && <p>Loading...</p>}
      {sessionError && <p>{sessionError.message}</p>}
      {session && <p>{session.ID}</p>}
    </div>
    <div>
      <UserButton />
    </div>
    </>
  
  );
}

export default Home;    