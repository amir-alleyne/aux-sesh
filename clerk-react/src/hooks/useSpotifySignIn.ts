import { useApi } from './useApi';
import {SpotifySignInResponse, spotifySignIn  } from '../api/spotify-auth';
import { useCallback } from 'react';

export function useSpotifySignIn() {
  // The generic hook infers the types for data and arguments.
  const { data, loading, error, callApi } = useApi<SpotifySignInResponse, []>(spotifySignIn);

  // Optionally, wrap callApi to create a clearer API for the hook user.
  const signInWithSpotify = useCallback(() => callApi(), [callApi]);


  return { response: data, loading, error, signInWithSpotify };
}