import { Input } from "@mui/material";
import { useEffect, useState } from "react";
import SpotifyWebApi from "spotify-web-api-node"
import TrackSearchResult, { Track } from "./TrackSearchResult";


const spotifyApi = new SpotifyWebApi({
    clientId: import.meta.env.VITE_SPOTIFY_CLIENT_ID,
  })

export function Session({token}: {token: string}) {
    const [searchResults, setSearchResults] = useState<any>([])
    const [search, setSearch] = useState<string>("");
 

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
        <Input placeholder="Search for a song" onChange={(e) => {
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
        </>
    )
}