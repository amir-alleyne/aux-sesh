  export interface SpotifySignInResponse {
    isSignedIn: boolean;
    // Include other response properties as needed
  }
  
  const SPOTIFY_CLIENT_ID = import.meta.env.VITE_SPOTIFY_CLIENT_ID

  export async function spotifySignIn(): Promise<SpotifySignInResponse> {
    // Build the Spotify authorization URL
    const clientId = SPOTIFY_CLIENT_ID;
    const redirectUri = encodeURIComponent('http://localhost:8080/auth-callback');
    const responseType = 'code';
    const scope = encodeURIComponent(
      'user-read-playback-state user-read-private user-read-email user-modify-playback-state'
    );
    const state = 'some-random-string';
    const spotifyAuthUrl = `https://accounts.spotify.com/authorize?client_id=${clientId}&redirect_uri=${redirectUri}&response_type=${responseType}&scope=${scope}&state=${state}`;
  
    // Calculate dimensions to center the popup
    const width = 600;
    const height = 800;
    const left = window.screenX + (window.innerWidth - width) / 2;
    const top = window.screenY + (window.innerHeight - height) / 2;
  
    // Open the popup window
    const popup = window.open(
      spotifyAuthUrl,
      'SpotifySignIn',
      `width=${width},height=${height},left=${left},top=${top}`
    );
  
    if (!popup) {
      throw new Error('Failed to open popup window');
    }
  
    // Return a promise that resolves when the popup is closed
    return new Promise<SpotifySignInResponse>((resolve, reject) => {
      const pollTimer = window.setInterval(() => {
        if (popup.closed) {
          window.clearInterval(pollTimer);
          // In a real app you would handle the actual auth response here.
          // For now, we simulate a successful sign-in.
          resolve({ isSignedIn: true });
        }
      }, 500);
    });
  }