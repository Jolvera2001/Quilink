import { useState } from "react"
import LoginForm from "../components/loginForm"
import { Button } from "@/components/ui/button"

function AuthPage() {
    const [isLoggingIn, toggleLogIn] = useState<boolean>(true)

    const handleToggle = () => {
        toggleLogIn(!isLoggingIn)
    }

    return(
        <>
            <div>
                {isLoggingIn
                    ? <LoginForm />
                    : <p>Register Form Here</p>
                }
                <Button onClick={handleToggle}>
                    {isLoggingIn ? "Need an Account? Register" : "Already have an account? Login"}
                </Button>
            </div>
        </>
    )
}

export default AuthPage