import {Card} from 'antd';

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
        <Card>
          <img src={track.albumUrl} style={{ height: "64px", width: "64px" }} />
          <div className="ml-3">
            <div>{track.title}</div>
            <div className="text-muted">{track.artist}</div>
          </div>
        </Card>
         
      )
}

export default TrackSearchResult;