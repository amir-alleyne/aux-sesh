import { BACKEND_BASE_URL } from "../hooks/useApi";

export interface SessionData {
    userId: string;
    // Include other fields as needed
  }
  
  export interface CreateSessionResponse {
    sessionId: string;
    // Include other response properties as needed
  }
  
  export async function createSession(sessionData: SessionData): Promise<CreateSessionResponse> {
    const response = await fetch(`${BACKEND_BASE_URL}/create-session`, {
      method: 'GET',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(sessionData),
    });
  
    if (!response.ok) {
      throw new Error('Error creating session');
    }
    return response.json();
  }