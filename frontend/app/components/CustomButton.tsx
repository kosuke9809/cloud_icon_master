'use client'
import React from "react"
import type { CustomButtonProps } from "@/types"
import { Button } from "@chakra-ui/react"
import { color } from "framer-motion"


export default function CustomButton({
    children,
    ...props
} : 
    CustomButtonProps
) {
    return (
        <Button 
            {...props}
        >
            {children}
        </Button>
    )
}