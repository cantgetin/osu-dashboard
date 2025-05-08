import React from 'react';
import {Line} from "react-chartjs-2";
import 'chart.js/auto';
import type {ChartOptions} from "chart.js";
import aveta from "aveta";
import {MakeChartHeightPlugin} from "./UserDiagrams.tsx";

interface LineChartProps {
    data: any;
    options: ChartOptions<'line'>;
}

const gridColor = 'rgba(115,115,115,0.2)';

export function generateOptions(titleText: string): ChartOptions<'line'> {
    return {
        aspectRatio: 2.5,
        plugins: {
            legend: {
                position: 'top',
                display: false,
            },
            title: {
                align: 'center',
                display: true,
                text: titleText,
                position: 'top',
                font: {
                    size: 14,
                    style: 'normal',
                    weight: 'normal',
                    family: 'Roboto',
                },
                color: 'rgb(255,255,255)',
            },
        },
        interaction: {
            intersect: false,
            mode: 'index',
        },
        scales: {
            x: {
                border: {
                    display: true
                },
                grid: {
                    display: true,
                    drawOnChartArea: true,
                    drawTicks: true,
                    color: gridColor,
                }
            },
            y: {
                grid: {
                    display: true,
                    drawOnChartArea: true,
                    drawTicks: true,
                    color: gridColor,
                },
                ticks: {
                    callback: (value) => {
                        return aveta(Number(value))
                    }
                }
            }
        }
    };
}


const LineChart: React.FC<LineChartProps> = ({data, options}) => {
    return <Line plugins={[MakeChartHeightPlugin(10)]} data={data} options={options}/>;
};

export default LineChart;