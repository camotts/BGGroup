import useUser from 'lib/useUser'
import Layout from 'components/Layout'

const Test = () => {

    const { user } = useUser({ redirectTo: '/login' })

    if (!user || user.isLoggedIn === false) {
        return <Layout>Loading...</Layout>
    }
}

export default Test