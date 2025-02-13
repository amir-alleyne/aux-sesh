import { UserButton, useAuth, useUser } from "@clerk/clerk-react";
import Button from '@mui/material/Button';
import { useSpotifySignIn } from "../hooks/useSpotifySignIn";
import { useCreateSession } from "../hooks/useCreateSession";
import { useQueueSong } from "../hooks/useQueueSong";
import { useEffect, useState } from "react";
import { Session } from "./Session";

function Home() {

  const { response, loading: signInLoading, error:signInError, signInWithSpotify } = useSpotifySignIn();
  const { session, loading: sessionLoading, error: sessionError, createAuxSession } = useCreateSession();
  const { status, loading: queueLoading, error: queueError, queueSongWithParams } = useQueueSong();
  const { isSignedIn, user, isLoaded } = useUser();
  useEffect(() => {
    if (user?.publicMetadata['spotify_token']) {
        var stringToken = user?.publicMetadata['spotify_token'] as string
        setToken(stringToken)
    }
}
, [user]);
  const {getToken} = useAuth();
  const [token, setToken] = useState<String | null>(null);

  useEffect(() => {
    console.log(token);
  }
  , [token]);


  const handleSpotifySignIn = async () => {
      console.log('Sign In with Spotify'); 
      try {
          await signInWithSpotify();
          setToken(await getToken());
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

    const handleQueueSong = async () => {
      try {
        if (token) {
           await queueSongWithParams({song_id: "6rqhFgbbKwnb9MLmUQDhG6", session_id: '1739408337', token: token as string});
        }
    } catch (err) {
        console.error('Queue song failed', err);
    }
    };
    
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
        <Button variant="outlined" onClick={handleQueueSong}>Queue Song</Button>
      </div>
      <div>
        {queueLoading && <p>Loading...</p>}
        {queueError && <p>{queueError.message}</p>}
        {status && <p>{status.Status}</p>}
      </div>
      {token && <div>
        <Session token={token as string}/>
      </div>}
      
      <div>
        <UserButton />
      </div>
    </>
  
  );
}

export default Home;    