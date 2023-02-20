-- 平台统计-任务 table.任务提交情况
-- 空间名称，提交任务人数，提交任务次数，任务失败次数，任务数量，训练总时长，单任务最高训练时长，GPU总卡时，单任务最高GPU卡时
SELECT
    a.space_id as space_id,
    b.space_name as space_name,
    c.commit_user_count as commit_user_count,
    SUM(a.is_first_commit) as commit_count,
    SUM(a.is_first_fail) as fail_count,
    COUNT(*) as job_count,
    SUM(a.job_duration) as total_job_duration,
    MAX(a.job_duration) as max_job_duration,
    SUM(a.gpu_duration) as total_gpu_duration,
    MAX(a.gpu_duration) as max_gpu_duration
FROM chart_job_day a
LEFT JOIN pub_space b ON a.space_id = b.space_id
LEFT JOIN (
    SELECT
        space_id,
        COUNT(*) as commit_user_count
    FROM (
         SELECT
             space_id, create_user_id
         FROM chart_job_day
         WHERE sample_timestamp = '"+ sample_timestamp +"' AND is_first_commit = 1
         GROUP BY space_id, create_user_id
    ) d
    GROUP BY d.space_id
) as c ON a.space_id = c.space_id
WHERE a.sample_timestamp = '"+ sample_timestamp +"'
GROUP BY a.space_id
WITH ROLLUP
ORDER BY a.space_id
LIMIT offset, count;

-- 空间统计页面 table.提交任务统计
-- 用户名，姓名，提交任务次数，任务数量，GPU总卡时，单任务最高GPU卡时，单任务平均卡时，训练总时长，单任务最高训练时长，单任务平均训练时长

