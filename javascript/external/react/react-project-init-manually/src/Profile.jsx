import { useParams } from 'react-router-dom';

const data = {
    memberA: {
        name: "Alice",
        desc: "the developer",
    },
    memberB: {
        name: "Bob",
        desc: "the designer",
    }
};


const Profile = () => {
    const params = useParams();
    const profile = data[params.username];

    return(
        <div>
            <h1>user profile</h1>
            {profile? (
                <div>
                    <h2>{profile.name}</h2>
                    <p>{profile.desc}</p>
                </div>                
            ):(
                <p>not exists</p>
            )}
        </div>
    )
}

export default Profile;