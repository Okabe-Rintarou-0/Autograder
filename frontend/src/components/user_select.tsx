import { useDebounceFn } from "ahooks";
import { Select, Spin } from "antd";
import { useState } from "react";
import { PAGE_SIZE } from "../lib/config";
import { User } from "../model/user";
import { listUsers } from "../service/user";

interface UserSelectProps {
    placeHolder?: string
    onChange: (userID: number) => void;
}

export default function UserSelect(props: UserSelectProps) {
    const { onChange, placeHolder } = props;
    const [fetching, setFetching] = useState<boolean>(false);
    const [users, setUsers] = useState<User[]>([]);
    const [loadingMore, setLoadingMore] = useState(false);
    const [userPageNo, setUserPageNo] = useState<number>(1);
    const [userKeyword, setUserKeyword] = useState<string>("");
    const [userHasNextPage, setUserHasNextPage] = useState<boolean>(false);
    const { run: debounceFetcher } = useDebounceFn(
        (keyword: string) => {
            keyword = keyword.trim();
            if (keyword) {
                setFetching(true);
                setUserKeyword(keyword);
                setUserPageNo(1);
                listUsers(keyword, 1, PAGE_SIZE).then((resp) => {
                    if (resp.data) {
                        const total = (userPageNo - 1) * PAGE_SIZE + resp.data.length;
                        setUserHasNextPage(total < resp.total);
                        setUsers(resp.data);
                        setUserPageNo(userPageNo => userPageNo + 1);
                    } else {
                        setUsers([]);
                        setUserHasNextPage(false);
                    }
                    setFetching(false);
                });
            }
        },
        { wait: 400 }
    );

    const loadMore = () => {
        if (!userHasNextPage) {
            return;
        }
        if (userKeyword) {
            setLoadingMore(true);
            listUsers(userKeyword, userPageNo, PAGE_SIZE).then((resp) => {
                if (resp.data) {
                    const total = (userPageNo - 1) * PAGE_SIZE + resp.data.length
                    setUserHasNextPage(total < resp.total);
                    setUsers(users.concat(resp.data));
                    setUserPageNo(userPageNo => userPageNo + 1);
                }
                setLoadingMore(false);
            });
        }
    }

    const onPopupScroll = (e: { target?: any }) => {
        const { target } = e;
        if (target.scrollTop + target.offsetHeight >= target.scrollHeight) {
            loadMore();
        }
    };

    return <Select
        style={{ width: "200px" }}
        allowClear
        showSearch
        onChange={onChange}
        placeholder={placeHolder ?? "指定用户"}
        filterOption={false}
        onSearch={debounceFetcher}
        onPopupScroll={onPopupScroll}
        notFoundContent={fetching ? <Spin size="small" /> : null}
        dropdownRender={(menu) => (
            <>
                {menu}
                {loadingMore ? (
                    <Spin size="small" style={{ textAlign: 'center' }} />
                ) : null}
            </>
        )}>
        {users.map((user) => {
            return <Select.Option key={user.id} value={user.id}>
                {`${user.real_name}(${user.username})`}
            </Select.Option>
        })}
    </Select>
}