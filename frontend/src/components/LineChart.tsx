import React from 'react';
import { Line } from "react-chartjs-2";
import 'chart.js/auto';
import type {ChartOptions} from "chart.js";

interface LineChartProps {
    chartData: any;
}

const gridColor = 'rgba(115,115,115,0.2)';


function formatNumber(str: string): string {
    const abbreviations = ["", "K", "M", "B", "T"];

    // Remove spaces and other non-numeric characters
    const cleanedStr = str.replace(/[^\d.]/g, '');

    if (!cleanedStr) {
        return '0';
    }

    const num = parseFloat(cleanedStr);
    const isNegative = num < 0;
    const absNum = Math.abs(num);

    const log1000 = Math.floor(Math.log10(absNum) / 3);
    const formattedNum = (absNum / Math.pow(1000, log1000)).toFixed(1);

    return (isNegative ? '-' : '') + formattedNum + abbreviations[log1000];
}
const options: ChartOptions<'line'> = {
    aspectRatio: 2.5,
    plugins:{
        legend: {
            position: 'top',
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
                    return formatNumber(value.toString())
                }
            }
        }
    }
}

const LineChart: React.FC<LineChartProps> = ({ chartData }) => {
    return <Line data={chartData} options={options} />;
};

export default LineChart;