import Layout from "../components/ui/Layout.tsx";
import Pagination from "../components/ui/Pagination.tsx";
import List from "../components/logic/List.tsx";
import Container from "../components/ui/Container.tsx";
import {useAppDispatch, useAppSelector} from "../store/hooks.ts";
import {fetchLogs, fetchLogsProps, LogsState, selectLogs} from "../store/logsSlice.ts";
import {useEffect} from "react";
import {LoadingState} from "../interfaces/LoadingState.ts";
import LoadingSpinner from "../components/ui/LoadingSpinner.tsx";
import ProgressBar from "../components/ui/ProgressBar.tsx";
import {CalendarIcon, ClockIcon, ServerIcon} from "../components/features/log/Icons.tsx";
import {formatDate, formatDuration} from "../utils/time.ts";
import {formatNanosToMilliseconds} from "../utils/utils.ts";
import {TagPill} from "../components/features/log/TagPill.tsx";

// TODO: refactor this, split into components
const Logs = ({...props}: fetchLogsProps) => {
    const dispatch = useAppDispatch();

    const logsState = useAppSelector<LogsState>(selectLogs)

    useEffect(() => {
        dispatch(fetchLogs(props))
    }, [dispatch])

    const onPageChange = (page: number) => {
        dispatch(fetchLogs({...props, page: page} as fetchLogsProps))
    }

    return (
        <Layout className="flex justify-center" title="logs">
            <Container>
                {
                    logsState.loading == LoadingState.Succeeded ?
                        <>
                            <List
                                className="w-full px-2 sm:px-0 grid grid-cols-1 gap-4"
                                items={logsState.logs!}
                                renderItem={(log: Log) => (
                                    <div
                                        key={log.id}
                                        className="flex flex-col sm:flex-row gap-4 p-4 bg-zinc-800 bg-opacity-30 rounded-lg
                                        hover:bg-zinc-700 transition-colors w-full"
                                    >
                                        <div className="flex-1 flex flex-col gap-2">
                                            <div className="flex items-center justify-between">
                                                <h3 className="text-lg font-semibold text-white">{log.name}</h3>
                                                <span className="text-sm text-zinc-400">ID: {log.id}</span>
                                            </div>

                                            {/* Message section */}
                                            {log.message && (
                                                <div className="flex items-start gap-1 text-sm">
                                                    <span className="font-medium text-zinc-400 flex-1 ml-1">
                                                        {log.message}
                                                    </span>
                                                </div>
                                            )}

                                            <div className="flex flex-wrap gap-4 text-sm">
                                                <div className="flex items-center gap-1">
                                                    <ClockIcon className="w-4 h-4 text-blue-400"/>
                                                    <span className="text-zinc-300">Duration:</span>
                                                    <span className="font-medium text-white">
                                                        {formatDuration(log.elapsed_time!)}
                                                    </span>
                                                </div>

                                                <div className="flex items-center gap-1">
                                                    <ServerIcon className="w-4 h-4 text-purple-400"/>
                                                    <span className="text-zinc-300">API Requests:</span>
                                                    <span className="font-medium text-white">
                                                        {log.api_requests!.toLocaleString()}
                                                    </span>
                                                </div>

                                                <div className="flex items-center gap-1">
                                                    <CalendarIcon className="w-4 h-4 text-green-400"/>
                                                    <span className="text-zinc-300">Finished at:</span>
                                                    <span className="font-medium text-white">
                                                        {formatDate(log.tracked_at!)}
                                                    </span>
                                                </div>

                                                {log.time_since_last_track && (
                                                    <div className="flex items-center gap-1">
                                                        <ClockIcon className="w-4 h-4 text-blue-400"/>
                                                        <span className="text-zinc-300">Since last track:</span>
                                                        <span className="font-medium text-white">
                                                            {formatDuration(log.time_since_last_track)}
                                                        </span>
                                                    </div>
                                                )}
                                            </div>

                                            <div className="mt-2 flex flex-wrap gap-3">
                                                <TagPill color="blue" text={log.platform!}/>
                                                <TagPill color="green" text={log.app_version!}/>
                                                {log.type && <TagPill color="purple" text={log.type}/>}
                                                {log.service && <TagPill color="amber" text={log.service}/>}
                                            </div>
                                        </div>

                                        <div className="w-full sm:w-[300px] sm:min-w-[300px] flex flex-col gap-2">
                                            <div className="flex justify-between text-sm">
                                                <span className="text-zinc-400">Success Rate:</span>
                                                <span className="font-medium text-white">
                                                    {log.success_rate_percent}%
                                                </span>
                                            </div>
                                            <ProgressBar value={log.success_rate_percent!} max={100} className="h-2"/>

                                            <div className="flex justify-between text-sm mt-2">
                                                <span className="text-zinc-400">Avg. Response Time:</span>
                                                <span className="font-medium text-white">
                                                    {formatNanosToMilliseconds(log.avg_response_time!)}
                                                </span>
                                            </div>
                                            <ProgressBar
                                                value={log.avg_response_time! / 1e6} // Convert ns to ms for progress bar
                                                max={1000}
                                                className="h-2"
                                                colorClass="bg-yellow-500"
                                            />
                                        </div>
                                    </div>
                                )}
                            />

                            <Pagination
                                pages={logsState.pages}
                                currentPage={logsState.currentPage}
                                onPageChange={onPageChange}
                                className="flex gap-1 md:gap-2 justify-end text-sm md:text-md"
                            />
                        </>
                        : <LoadingSpinner/>
                }
            </Container>
        </Layout>
    );
};

export default Logs;