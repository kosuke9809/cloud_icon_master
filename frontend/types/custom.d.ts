import React from "react"
import { Button } from "@chakra-ui/react"
import { jsx } from "@emotion/react";

type ButtonProps = React.ComponentProps<typeof Button>

type ButtonColorScheme = "whiteAlpha" | "blackAlpha" | "gray" | "red" | "orange" | "yellow" | "green" | "teal" | "blue" | "cyan" | "purple" | "pink" | "linkedin" | "facebook" | "messenger" | "whatsapp" | "twitter" | "telegram";

interface CustomButtonProps extends ButtonProps {
    children?: React.ReactNode;
    onClick?: React.MouseEventHandler<HTMLButtonElement>;
    colorScheme?: ButtonColorScheme;
    variant?: "ghost" | "outline" | "solid" | "link" | "unstyled";
    iconSpacing?: ResponsiveValue<string | number | (string & {})>;
    isActive?: boolean;
    isDisabled?: boolean;
    isLoading?: boolean;
    leftIcon?: ReactElement<any, string | JSXElementConstructor<any>>;
    rightIcon?: ReactElement<any, string | JSXElementConstructor<any>>;
    size?: "lg" | "md" | "sm" | "xs";
    spinner?: ReactElement<any, string | JSXElementConstructor<any>>; 
    spinnerPlacement?: "start" | "end";
}

type Quiz = {
    id: number;
    name: string;
    category: string;
    object_key: string;
}

type QuizData = Quiz[]

type QuizProps = {
    iconImgUrl : string
    answer: string
    quizData:  QuizData
}

type QuizResult = {
    answer: string | null
    select: string | null,
    result: boolean,
}

type State = {
    AllQuizResult: QuizResult[]
    CorrectNum: number
    countCorrectNum: () => void
    addQuizResult: (payload: QuizResult) => void
    resetAllQuizResult: () => void
    resetCorrectNum: () => void
}