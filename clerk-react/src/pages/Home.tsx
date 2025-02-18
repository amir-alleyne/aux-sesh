import { UserButton, useAuth } from "@clerk/clerk-react";
import { useSpotifySignIn } from "../hooks/useSpotifySignIn";
import { useEffect, useState } from "react";
import GenModal from "../components/modal";
import { Button, Input } from "antd";
import { useJoinSession } from "../hooks/useJoinSession";
import { useLeaveSession } from "../hooks/useLeaveSession";



function Home() {
  const [openCreateSession, setOpenCreateSession] = useState<boolean>(false);
  const [openJoinSession, setOpenJoinSession] = useState<boolean>(false);
  const [sessionCreated, setSessionCreated] = useState<boolean>(false);
  const [sessionId, setSessionId] = useState<number>(0);
  const [openLeaveSession, setOpenLeaveSession] = useState<boolean>(false);
  const { response, loading: signInLoading, error:signInError, signInWithSpotify } = useSpotifySignIn();
  const { joinResponse , loading: isJoining, error: joinError, joinAuxSession } =  useJoinSession();
  const { leaveResponse , loading: isLeaving, error: leaveError, leaveAuxSession } =  useLeaveSession();

  // const { session, loading: sessionLoading, error: sessionError, createAuxSession } = useCreateSession();
 
  const {getToken} = useAuth();
  const [token, setToken] = useState<String | null>(null);

  useEffect(() => {
    const fetchToken = async () => {
      const token = await getToken();
      setToken(token);
    }
    fetchToken();
  }
  , [token]);


  const handleSpotifySignIn = async () => {
      console.log('Sign In with Spotify'); 
      try {
          await signInWithSpotify();
          
          // setToken(await getToken());
      } catch (err) {
          console.error('Sign In failed:', err);
      }
      };

  const handleJoinSession = async (sessionId: number) => {
      
        try {
            await joinAuxSession(sessionId, token as string);
            setOpenJoinSession(false);
        } catch (err) {
            console.error('Create Session failed:', err);
        }
        }
  const handleLeaveSession = async (sessionId: number) => {
      
        try {
            await leaveAuxSession(sessionId, token as string);
            setOpenLeaveSession(false);
        } catch (err) {
            console.error('Create Session failed:', err);
        }
  }
  const handleCreateSession = () => {
    
      // try {
      //     await createAuxSession();
      // } catch (err) {
      //     console.error('Create Session failed:', err);
      // }
      }

 
  return (
    <>
      
      <div>
        <Button variant="outlined" onClick={() => {
          setOpenCreateSession(true);
        }}>Create Session</Button>
      </div>
      {openCreateSession && <div>
        <GenModal title="Create Session" open={openCreateSession} onClose={() => setOpenCreateSession(false)}>
          <div>
            {sessionCreated ?  
            <p>Session Created</p>

            :
            <Button variant="outlined" onClick={handleSpotifySignIn}>Sign In With Spotify</Button>
            }
            
          </div>
        </GenModal>
      </div>}

      <div>
        <Button variant="outlined" onClick={() => {
          setOpenJoinSession(true);
        }}>Join Session</Button>
      </div>
      {openJoinSession && <div>
        <GenModal title="Join Session" open={openJoinSession} onClose={() => setOpenJoinSession(false)} onOk={() => handleJoinSession(sessionId)}>
          <div>
            <Input type="number" placeholder="Enter Session ID" onChange={(e) => { setSessionId(Number(e.target.value))}} />
            
          </div>
        </GenModal>
      </div>}

      <div>
        <Button variant="outlined" onClick={() => {
          setOpenLeaveSession(true);
        }}>Leave Session</Button>
      </div>
      {openLeaveSession && <div>
        <GenModal title="Leave Session" open={openLeaveSession} onClose={() => setOpenLeaveSession(false)} onOk={() => handleLeaveSession(sessionId)}>
          <div>
            <Input type="number" placeholder="Enter Session ID" onChange={(e) => { setSessionId(Number(e.target.value))}} />
            
          </div>
        </GenModal>
      </div>}
  
     

      
    
    </>
  
  );
}

export default Home;