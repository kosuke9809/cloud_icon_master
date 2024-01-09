"use client"
import { Box, Flex, Text, Avatar  } from "@chakra-ui/react";
import { Link } from "@chakra-ui/next-js";

export default function Navbar() {
    return (
        <Box bg="#232F3E" w="100%" p={4} color="white">
        <Flex alignItems="center" justifyContent="space-between">
            <Link  px={2} href="/">
                <Text color="#FF9900" fontSize="lg" fontWeight="bold" fontFamily="-moz-initial">AWS ICON MASTER</Text>
            </Link>
            <Flex>
            <Link px={2} href="/mypage">
                <Avatar  size="sm" name="None" src="https://menslog.net/wp-content/uploads/2021/03/%E3%82%BB%E3%83%B3%E3%83%AA%E3%83%84%EF%BC%95-790x480.jpg"/>
            </Link>
            </Flex>
        </Flex>
        </Box>
    )
}