import { useApi } from './useApi';
import { useCallback } from 'react';
import { QueueSongParams, QueueSongResponse, queueSong } from '../api/sessions';


export function useQueueSong() {
  // The generic hook infers the types for data and arguments.
  const { data, loading, error, callApi } = useApi<QueueSongResponse, [QueueSongParams]>(queueSong);

  // Optionally, wrap callApi to create a clearer API for the hook user.
  const queueSongWithParams = useCallback((params: QueueSongParams) => callApi(params), [callApi]);


  return { status: data, loading, error, queueSongWithParams };
}