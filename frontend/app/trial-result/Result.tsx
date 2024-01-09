"use client"

import useStore from "@/store"
import { Text, Card, CardHeader, CardBody, CardFooter, Heading } from "@chakra-ui/react"
import CustomButton from "../components/CustomButton"
import type { CustomButtonProps } from "@/types"

export default function Result () {
    const {AllQuizResult, CorrectNum} = useStore()

    const loginButtonProps: CustomButtonProps = {
        colorScheme: "orange",
        variant: "outline",
    }

    return (
        <Card align="center" className=" mt-10">
            <CardHeader>
                <Heading size={"md"}>結果</Heading>
            </CardHeader>
            <CardBody>
                <Text>正解数：{CorrectNum} / 5</Text>
                <Text>間違えたサービス</Text>
                {
                    AllQuizResult.map((quiz, index) => (
                        quiz.result ? null : 
                        <Text key={index}>{quiz.answer}</Text>
                    ))
                }
            </CardBody>
            <CardFooter>
                <CustomButton {...loginButtonProps} >ログイン</CustomButton>
            </CardFooter>
        </Card>
    )
}