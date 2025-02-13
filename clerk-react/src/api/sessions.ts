import { useAuth } from "@clerk/clerk-react";
import { BACKEND_BASE_URL } from "../hooks/useApi";

export interface SessionData {
    userId: string;
    // Include other fields as needed
  }
  
  export interface CreateSessionResponse {
    ID: string;
    AdminID: string;
    UserIDs: string[];
    SongQueue: string[];
  }
  
  export interface QueueSongResponse {
    Status: boolean;
  }

  export interface QueueSongParams {
    song_id: string;
    session_id: string;
    token: string;
  }
  
  export async function createSession({token} : {token: string}): Promise<CreateSessionResponse> {
    const response = await fetch(`${BACKEND_BASE_URL}/create-session`, {
      method: 'GET',
      headers: { 'Content-Type': 'application/json',
      "Authorization": `Bearer ${token}`
      },
    });
  
    if (!response.ok) {
      throw new Error('Error creating session');
    }
    return response.json();
  }

  export async function queueSong({song_id, session_id, token}: QueueSongParams): Promise<QueueSongResponse> {
    
    const response = await fetch(`${BACKEND_BASE_URL}/queue-song`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json',
      "Authorization": `Bearer ${token}`
      },
      body: JSON.stringify({
        "song_id": song_id,
        "session_id": session_id
      })
    });
  
    if (!response.ok) {
      throw new Error('Error creating session');
    }
    return response.json();
  }