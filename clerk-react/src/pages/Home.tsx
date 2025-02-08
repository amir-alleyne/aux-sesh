import { UserButton } from "@clerk/clerk-react";
import Button from '@mui/material/Button';
import { useSpotifySignIn } from "../hooks/useSpotifySignIn";


function Home() {

const { response , loading, error, signInWithSpotify } = useSpotifySignIn();

const handleSpotifySignIn = async () => {
    console.log('Sign In with Spotify');
    try {
        await signInWithSpotify();
    } catch (err) {
        console.error('Sign In failed:', err);
    }
    };

  return (
    <>
      <div>
      <Button variant="outlined" onClick={handleSpotifySignIn}>Sign In With Spotify</Button>
    </div>
    <div>
      {loading && <p>Loading...</p>}
      {error && <p>{error.message}</p>}
      {response && <p>{response.isSignedIn}</p>}
    </div>
    <div>
      <Button variant="outlined" onClick={handleSpotifySignIn}>Create Session</Button>
    </div>
    <div>
      <UserButton />
    </div>
    </>
  
  );
}

export default Home;    