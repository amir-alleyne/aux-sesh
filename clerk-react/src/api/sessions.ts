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
  
  export async function createSession(): Promise<CreateSessionResponse> {
    const response = await fetch(`${BACKEND_BASE_URL}/create-session`, {
      method: 'GET',
      headers: { 'Content-Type': 'application/json' },
    });
  
    if (!response.ok) {
      throw new Error('Error creating session');
    }
    return response.json();
  }