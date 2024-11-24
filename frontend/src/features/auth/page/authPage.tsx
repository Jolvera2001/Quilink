import { useState } from "react"
import LoginForm from "../components/loginForm"
import { Button } from "@/components/ui/button"
import RegisterForm from "../components/registerForm"
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card"

function AuthPage() {
    const [isLoggingIn, toggleLogIn] = useState<boolean>(true)

    const handleToggle = () => {
        toggleLogIn(!isLoggingIn)
    }

    return(
        <div className="h-screen flex items-center justify-center">
                <Card className="w-[350px]">
                    <CardHeader>
                        <CardTitle>{isLoggingIn ? "Log In" : "Register"}</CardTitle>
                        <CardDescription>
                            {isLoggingIn
                                ? "Log into an existing account"
                                : "Create a new account"
                            }
                        </CardDescription>
                    </CardHeader>
                    <CardContent className="space-y-4">
                        {isLoggingIn
                            ? <LoginForm />
                            : <RegisterForm />
                        }
                        <Button onClick={handleToggle}>
                            {isLoggingIn ? "Need an Account? Register" : "Already have an account? Login"}
                        </Button>
                    </CardContent>
                </Card>
        </div>
    )
}

export default AuthPage