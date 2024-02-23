import LineChart from "./LineChart.tsx";
import {convertDataToDayMonth} from "../utils/utils.ts";
import {useEffect, useState} from "react";

interface LineChartProps {
    data: UserStatsDataset[]
    asSlideshow?: boolean
}

const UserCharts = (props: LineChartProps) => {

    function mapToChartData(data: UserStatsDataset[]): UserStatsDataset[] {
        const updatedArray: UserStatsDataset[] = [];
        data.forEach((obj) => updatedArray.push({...obj, timestamp: convertDataToDayMonth(obj.timestamp)}));
        return updatedArray
    }

    type ChartDataProperty = keyof UserStatsDataset;

    const generateSingleChartData = (
        userData: UserStatsDataset[],
        property: ChartDataProperty,
        name: string,
        color: string
    ) => {
        if (userData.length >= 7) {
            userData = userData.slice(-7)
        }

        return {
            labels: userData.map((data) => data.timestamp),
            datasets: [{
                data: userData.map((data) => data[property]),
                backgroundColor: [color],
                label: name,
                borderColor: [color],
                borderWidth: 2,
                pointStyle: 'circle',
                pointRadius: 3,
                pointHoverRadius: 6,
            }],
        };
    };

    const playCountChart = () => {
        return (
            <LineChart
                chartData={
                    generateSingleChartData(
                        mapToChartData(props.data),
                        'play_count',
                        'Play count',
                        '#86EFAC'
                    )
                }
            />
        )
    }

    const favouritesChart = () => {
        return (
            <LineChart
                chartData={
                    generateSingleChartData(
                        mapToChartData(props.data),
                        'favourite_count',
                        'Favourite count',
                        '#FF5DBD'
                    )
                }
            />
        )
    }

    const mapCountChart = () => {
        return (
            <LineChart
                chartData={
                    generateSingleChartData(
                        mapToChartData(props.data),
                        'map_count',
                        'Map count',
                        '#ffed54'
                    )
                }
            />
        )
    }

    const commentsChart = () => {
        return (
            <LineChart
                chartData={
                    generateSingleChartData(
                        mapToChartData(props.data),
                        'comments_count',
                        'Comment count',
                        '#f87171'
                    )
                }
            />
        )
    }

    const chartsList = [playCountChart, favouritesChart, mapCountChart, commentsChart]

    const intervalInSeconds = 3;
    const [currentIndex, setCurrentIndex] = useState(0);

    useEffect(() => {
        const intervalId = setInterval(() => {
            setCurrentIndex((prevIndex) => (prevIndex + 1) % chartsList.length);
        }, intervalInSeconds * 1000);

        return () => clearInterval(intervalId);
    }, [chartsList.length, intervalInSeconds]);

    return (
        <>
            {
                props.data.length > 0 ?
                    <div className="flex gap-3 bg-zinc-900 rounded-lg p-2 box-border w-full">
                        {props.asSlideshow ?
                            <>{chartsList[currentIndex]()}</>
                            :
                            <div className="grid grid-cols-2 w-full">
                                <div className="w-full">
                                    {playCountChart()}
                                </div>
                                <div>
                                    {favouritesChart()}
                                </div>
                                <div>
                                    {mapCountChart()}
                                </div>
                                <div>
                                    {commentsChart()}
                                </div>
                            </div>
                        }
                    </div>
                    : null}
        </>
    );
};

export default UserCharts;