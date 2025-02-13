export type Track = {
    id: string;
    title: string;
    artist: string[];
    albumUrl: string;
    image: string;
    uri: string;
    };



function TrackSearchResult({ track }: { track: Track}) {

    return (
        <div>
          <img src={track.albumUrl} style={{ height: "64px", width: "64px" }} />
          <div className="ml-3">
            <div>{track.title}</div>
            <div className="text-muted">{track.artist}</div>
          </div>
        </div>
      )
}

export default TrackSearchResult;