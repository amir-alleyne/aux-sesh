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
  
    
  return new Promise<SpotifySignInResponse>((resolve, reject) => {
    let resolved = false;

    // Listen for the message from the backend callback page.
    const messageHandler = (event: MessageEvent) => {
      // Verify that the message is coming from the expected origin
      if (event.origin !== window.location.origin) {
        return;
      }
      const data = event.data;
      if (data && data.type === 'spotify-auth-callback') {
        resolved = true;
        window.removeEventListener('message', messageHandler);
        clearInterval(pollTimer);

        // Resolve based on the response from your backend
       if (data.isSignedIn) {
          resolve({ isSignedIn: true });
        } else {
          resolve({ isSignedIn: false });
        }

        if (!popup.closed) {
          popup.close();
        }
      }
    };

    window.addEventListener('message', messageHandler);

    // Poll for popup closure: if the user manually closes the popup before authentication completes
    const pollTimer = window.setInterval(() => {
      if (popup.closed && !resolved) {
        window.removeEventListener('message', messageHandler);
        clearInterval(pollTimer);
        resolve({ isSignedIn: false });
      }
    }, 500);
  });
  }