import Layout from "../components/ui/Layout.tsx";
import UserList from "../components/features/user/UserList.tsx";

const Users = () => {
    return (
        <Layout className="flex justify-center" title="Users">
            <UserList
                page={1}
                sort="playcount"
                direction="desc"
            />
        </Layout>
    );
};

export default Users;