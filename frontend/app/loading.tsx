import { Spinner } from "@chakra-ui/react";


export default function Loading ()  {
    return (
        <div className="flex flex-1 mt-10 justify-center">
            <Spinner color="gray" size="lg" speed='0.65s' thickness='4px'/>
        </div>
    )
}