import { useEffect, useState } from "react";
import SpotifyWebApi from "spotify-web-api-node"
import TrackSearchResult, { Track } from "./TrackSearchResult";
import { useQueueSong } from "../hooks/useQueueSong";
import { Button, Input } from "antd";
import { useUser } from "@clerk/clerk-react";

const spotifyApi = new SpotifyWebApi({
    clientId: import.meta.env.VITE_SPOTIFY_CLIENT_ID,
  })

export function Session() {
    const [searchResults, setSearchResults] = useState<any>([])
    const [token, setToken] = useState<string>("");
    const { isSignedIn, user, isLoaded } = useUser();
    const [search, setSearch] = useState<string>("");
    const { status, loading: queueLoading, error: queueError, queueSongWithParams } = useQueueSong();
    const handleQueueSong = async () => {
      try {
        if (token) {
           await queueSongWithParams({song_id: "6rqhFgbbKwnb9MLmUQDhG6", session_id: '1739408337', token: token as string});
        }
    } catch (err) {
        console.error('Queue song failed', err);
    }
    };
    
 
    
    useEffect(() => {
      if (user?.publicMetadata['spotify_token']) {
          var stringToken = user?.publicMetadata['spotify_token'] as string
          setToken(stringToken)
      }
  }
  , [user]);

    useEffect(() => {
        if (!token) return
        spotifyApi.setAccessToken(token)
      }, [token])

    useEffect(() => {
        if (!search || !token) return setSearchResults([])
        let cancel = false
        spotifyApi.searchTracks(search).then((res) => {
            if (cancel) return
            setSearchResults(
                res?.body?.tracks?.items.map(track => {
                  const smallestAlbumImage = track.album.images.reduce(
                    (smallest, image) => {
                        if (image?.height !== undefined && (smallest.height === undefined || image.height < smallest.height)) {
                            return image;
                        }
                      return smallest
                    },
                    track.album.images[0]
                  )
        
                  return {
                    artist: track.artists[0].name,
                    title: track.name,
                    uri: track.uri,
                    albumUrl: smallestAlbumImage.url,
                  }
                })
              )
        })

    }, [search, token]);

    return (
        <>
        <div>
        <div>
        <Button variant="outlined" onClick={handleQueueSong}>Queue Song</Button>
        <Input  placeholder="Search for a song" onChange={(e) => {
            setSearch(e.target.value)
        }}/>
        </div>
        <div>
        {searchResults.map( (track: Track) => (
          <TrackSearchResult
            track={track}
            key={track.uri}
          />
        ))}
        </div>
        </div>
        </>
    )
}