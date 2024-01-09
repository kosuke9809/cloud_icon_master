import { GetObjectCommand, S3Client } from "@aws-sdk/client-s3"
import type { QuizProps } from "@/types"
import QuizPannel from "./QuizPannel"

async function fetchQuiz() {
    // await new Promise((resolve) => setTimeout(resolve, 2000))
    const res = await fetch (`${process.env.backendApiUrl}/quiz`,{
        cache: "no-store"
    })

    if (!res.ok) {
        throw new Error("Failed to fetch data in server")
    }

    const quiz = await res.json()
    return quiz
}

// S3から取得したストリームデータをバイト配列に変換する関数
async function streamToBuffer(stream:any) {
    const chunks = [];
    for await (const chunk of stream) {
        chunks.push(chunk);
    }
    return Buffer.concat(chunks);
}

// S3にプライベートアクセスして、アイコン画像を取得する関数
async function getIconFromS3(objectKey: string) {
    const credentials = {
        accessKeyId: process.env.AWS_ACCESS_KEY_ID,
        secretAccessKey: process.env.AWS_SECRET_ACCESS_KEY
    }
    const client = new S3Client({
        region: process.env.AWS_REGION,
        ...credentials
    })
    const params = {
        Bucket: process.env.AWS_S3_BUCKET_NAME,
        Key: objectKey
    }

    try{
        const command = new GetObjectCommand(params);
        const res = await client.send(command);
        const dataBuffer = await streamToBuffer(res.Body);
        const base64 = dataBuffer.toString('base64')
        const imageUrl = `data:image/jpeg;base64,${base64}`;
        return imageUrl
    } catch (error) {
        console.error("Error", error);
        throw error;
    }
}

export default async function Quiz() {
    const quiz = await fetchQuiz()
    const randomNum = Math.floor(Math.random() * 4);
    const ans = quiz[randomNum].name
    const iconImgUrl = await getIconFromS3(quiz[randomNum].object_key);
    const quizProps : QuizProps = {
        iconImgUrl: iconImgUrl,
        answer: ans,
        quizData: quiz
    }

    return (
        <div>
            <QuizPannel {...quizProps}></QuizPannel>
        </div>
    )
}